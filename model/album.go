package model


type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func New(id string, title string, artist string, price float64) Album{
	return Album{id, title, artist, price}
}

func GetAllAlbums() []Album{
	return albums
}

func GetAlbumById(id string) []Album{
	albums := albums
	for  _,row := range(albums){
		if row.ID == id{
			return []Album{row}
		}
	}
	return albums
}

func AddAlbum(album Album) []Album{
	albums = append(albums, album)	 
	return albums
}