package peliculas

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func obtenerDB() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreDB := "prueba"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreDB))

	if err != nil {
		return nil, err
	}
	return db, nil
}

func nuevaPeliculaDB(p Pelicula) (id int64, e error) {
	db, err := obtenerDB()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer db.Close()
	//PREPARAMOS PARA PREVENIR INYECC
	sentenciaPrep, err := db.Prepare("INSERT INTO peliculas (titulo, anio, idioma, sinopsis) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer sentenciaPrep.Close()
	//EJECUTAMOS
	i, err := sentenciaPrep.Exec(p.Titulo, p.Anio, p.Idioma, p.Sinopsis)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return i.LastInsertId()
}
func GetPeliculaDB(id int64) Pelicula {
	var p Pelicula
	db, err := obtenerDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fila, err := db.Prepare("SELECT titulo, anio, idioma, sinopsis, id FROM peliculas WHERE id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer fila.Close()
	err = fila.QueryRow(id).Scan(&p.Titulo, &p.Anio, &p.Idioma, &p.Sinopsis, &p.ID)
	if err != nil {
		fmt.Println(err)
	}
	return p
}
func ListaPeliculasDB() []Pelicula {
	var lista []Pelicula
	var p Pelicula
	db, err := obtenerDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//
	filas, err := db.Query("SELECT titulo,anio,idioma,sinopsis,id FROM peliculas")
	if err != nil {
		fmt.Println(err)
	}
	defer filas.Close()
	for filas.Next() {
		err = filas.Scan(&p.Titulo, &p.Anio, &p.Idioma, &p.Sinopsis, &p.ID)
		if err != nil {
			fmt.Println(err)
		}
		lista = append(lista, p)
	}
	return lista
}
func EliminarPeliculaDB(id int) []Pelicula {
	db, err := obtenerDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//
	sentenciaPrep, err := db.Prepare("DELETE FROM peliculas WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer sentenciaPrep.Close()
	//
	_, err = sentenciaPrep.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	return ListaPeliculasDB()
}

func ModificarPeliculaDB(pelicula Pelicula) []Pelicula {
	db, err := obtenerDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	sentenciaPrep, err := db.Prepare("UPDATE peliculas SET titulo = ?, anio = ?, idioma = ?, sinopsis = ? WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer sentenciaPrep.Close()
	_, err = sentenciaPrep.Exec(pelicula.Titulo, pelicula.Anio, pelicula.Idioma, pelicula.Sinopsis, pelicula.ID)
	if err != nil {
		fmt.Println(err)
	}
	return ListaPeliculasDB()
}
