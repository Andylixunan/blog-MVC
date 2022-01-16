package models

import "blogweb_gin/database"

type Image struct {
	ID         int
	FilePath   string
	FileName   string
	Status     int
	CreateTime int64
}

func InsertIntoAlbum(image Image) (int64, error) {
	return database.ModifyDB(
		"insert into album(filepath, filename, status, createTime) values(?, ?, ?, ?)",
		image.FilePath, image.FileName, image.Status, image.CreateTime,
	)
}

func FindAllImagesFromAlbum() ([]Image, error) {
	rows, err := database.QueryDB("select id, filepath, filename, status, createTime from album where status=0")
	if err != nil {
		return nil, err
	}
	images := []Image{}
	for rows.Next() {
		image := Image{}
		rows.Scan(&image.ID, &image.FilePath, &image.FileName, &image.Status, &image.CreateTime)
		images = append(images, image)
	}
	return images, nil
}
