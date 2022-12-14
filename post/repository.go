package post

var posts []*Post

func init() {
	posts = []*Post{
		{
			id:      1,
			title:   "post 1",
			content: "post 1",
		},
		{
			id:      2,
			title:   "post 2",
			content: "post 2",
		},
		{
			id:      3,
			title:   "post 3",
			content: "post 3",
		},
	}
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindByID(id int) (*Post, error) {
	for _, p := range posts {
		if p.ID() == id {
			return p, nil
		}
	}
	return nil, nil
}

func (r *Repository) FindAll() ([]*Post, error) {
	return posts, nil
}
