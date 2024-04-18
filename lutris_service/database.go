package lutris_service

import (
	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/dao"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_dao"
)

func UpsertLutrisDb(game *lutris_dao.LutrisGame, category *lutris_dao.LutrisCategory) error {
	// tx.begin
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		gameId     int64
		categoryId int64
	)

	// game
	err = tx.Get(&gameId, `--sql
	select id from games where slug = ?`, game.Slug)
	if err != nil {
		// insert
		res, err := tx.NamedExec(`--sql
		insert into games (
			name, slug, platform, lastplayed, playtime, runner, directory, installed, installed_at, updated, configpath, has_custom_banner, has_custom_icon, has_custom_coverart_big, hidden
		) values (
			:name, :slug, :platform, :lastplayed, :playtime, :runner, :directory, :installed, :installed_at, :updated, :configpath, :has_custom_banner, :has_custom_icon, :has_custom_coverart_big, :hidden
		)`, game)
		if err != nil {
			logger.Erro(err)
			return err
		}
		gameId, _ = res.LastInsertId()
	} else {
		// history
		if game.Lastplayed.Valid || game.Playtime.Valid {
			if _, err := tx.Exec(`--sql
			update games set lastplayed = ?, playtime = ? where id = ?`, game.Lastplayed, game.Playtime, gameId); err != nil {
				logger.Erro(err)
				return err
			}
		}
	}

	// categories
	err = tx.Get(&categoryId, `--sql
	select id from categories where name = ?`, category.Name)
	if err != nil {
		res, err := tx.Exec(`--sql
		insert into categories (name) values (?)`, category.Name)
		if err != nil {
			logger.Erro(err)
			return err
		}
		categoryId, _ = res.LastInsertId()
	}

	// rel
	_, err = tx.Exec(`--sql
	insert into games_categories (game_id, category_id) values (?, ?)`, gameId, categoryId)
	if err != nil {
		logger.Erro(err)
		return err
	}

	// tx.commit
	return tx.Commit()
}

func InstallLutrisDb(game *lutris_dao.LutrisGame, category *lutris_dao.LutrisCategory) error {
	// tx.begin
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		gameId     int64
		categoryId int64
	)

	// clean
	_, err = tx.Exec(`--sql
	delete from games_categories where game_id in (select id from games where slug = ?)`, game.Slug)
	if err != nil {
		logger.Warn(err)
	}
	_, err = tx.Exec(`--sql
	delete from categories where name = ?`, category.Name)
	if err != nil {
		logger.Warn(err)
	}
	_, err = tx.Exec(`--sql
	delete from games where slug = ?`, game.Slug)
	if err != nil {
		logger.Warn(err)
	}

	// game
	res, err := tx.NamedExec(`--sql
	insert into games (
		name, slug, platform, lastplayed, playtime, runner, directory, installed, installed_at, updated, configpath, has_custom_banner, has_custom_icon, has_custom_coverart_big, hidden
	) values (
		:name, :slug, :platform, :lastplayed, :playtime,:runner, :directory, :installed, :installed_at, :updated, :configpath, :has_custom_banner, :has_custom_icon, :has_custom_coverart_big, :hidden
	)`, game)
	if err != nil {
		logger.Erro(err)
		return err
	}
	gameId, _ = res.LastInsertId()

	// categories
	res, err = tx.Exec(`--sql
	insert into categories (name) values (?)`, category.Name)
	if err != nil {
		logger.Erro(err)
		return err
	}
	categoryId, _ = res.LastInsertId()

	// rel
	if _, err := tx.Exec(`--sql
	insert into games_categories (game_id, category_id) values (?, ?)`, gameId, categoryId); err != nil {
		logger.Erro(err)
		return err
	}

	// tx.commit
	return tx.Commit()
}

func CleanLutrisDb() error {
	// tx.begin
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		logger.Erro(err)
		return err
	}
	defer func() { _ = tx.Rollback() }()

	// rel
	if _, err := tx.Exec(`--sql
	delete from games_categories where game_id in (select id from games where slug like ?)`, config.SlugPrefix+"%"); err != nil {
		logger.Erro(err)
		return err
	}

	// games
	if _, err := tx.Exec(`--sql
	delete from games where slug like ?`, config.SlugPrefix+"%"); err != nil {
		logger.Erro(err)
		return err
	}

	// categories
	if _, err := tx.Exec(`--sql
	delete from categories where name like ? or name like ?`, "@%", "$%"); err != nil {
		logger.Erro(err)
		return err
	}

	// tx.commit
	return tx.Commit()
}

func ListNameAndLastplayedAndPlaytime() ([]*lutris_dao.LutrisGame, error) {
	db := dao.LuDb

	games := make([]*lutris_dao.LutrisGame, 0)
	err := db.Select(&games, `--sql
	SELECT name,
	slug,
    playtime,
    lastplayed
	from games
	where slug like ?
		and (
			playtime is not null
			or lastplayed is not null
		)
		and (
			abs(playtime - 0) > 0.01
			or lastplayed > 0
		)`, config.SlugPrefix+"%")
	if err != nil {
		return nil, err
	}

	return games, nil
}
