package lutris_service

import (
	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/dao"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/XDwanj/go-gsgm/lutris_dao"
)

func UpsertLutrisDb(
	game *lutris_dao.LutrisGame,
	category *lutris_dao.LutrisCategory,
) error {
	// tx.begin
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var (
		gameId     int64
		categoryId int64
	)

	// game
	err = tx.Get(&gameId, `--sql
	select * from games where slug = ?`, game.Slug)
	if err != nil {
		res, err := tx.NamedExec(`--sql
		insert into games (
			name, slug, platform, runner, directory, installed, installed_at, updated, configpath, has_custom_banner, has_custom_icon, has_custom_coverart_big, hidden
		) values (
			:name, :slug, :platform, :runner, :directory, :installed, :installed_at, :updated, :configpath, :has_custom_banner, :has_custom_icon, :has_custom_coverart_big, :hidden
		)`, game)
		if err != nil {
			logger.Erro(err)
			return err
		}
		gameId, _ = res.LastInsertId()
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

	_, err = tx.Exec(`--sql
	insert into games_categories (game_id, category_id) values (?, ?)`, gameId, categoryId)
	if err != nil {
		logger.Erro(err)
		return err
	}

	// tx.commit
	return tx.Commit()
}

func InstallLutrisDb(
	game *lutris_dao.LutrisGame,
	category *lutris_dao.LutrisCategory,
) error {
	// tx.begin
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var (
		gameId     int64
		categoryId int64
	)

	// clean
	tx.Exec(`--sql
	delete from games_categories where game_id in (select id from games where slug = ?)`, game.Slug)
	tx.Exec(`--sql
	delete from categories where id in (select id from categories where name = ?)`, category.Name)
	tx.Exec(`--sql
	delete from games where id in (select id from games where slug = ?)`, game.Slug)

	// game
	res, err := tx.NamedExec(`--sql
	insert into games (
		name, slug, platform, runner, directory, installed, installed_at, updated, configpath, has_custom_banner, has_custom_icon, has_custom_coverart_big, hidden
	) values (
		:name, :slug, :platform, :runner, :directory, :installed, :installed_at, :updated, :configpath, :has_custom_banner, :has_custom_icon, :has_custom_coverart_big, :hidden
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
	// tx
	tx, err := dao.LuDb.Beginx()
	if err != nil {
		logger.Erro(err)
		return err
	}
	defer tx.Rollback()

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
	tx.Exec(`--sql
	delete from categories where name like ? or name like ?`, "@%", "$%")

	// tx.commit
	return tx.Commit()
}
