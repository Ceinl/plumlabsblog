package articles

type Manager struct {
	article  Article
}

type Article struct {
	Title       string
	mdContent   string
	htmlContent string
}

func NewArticleManager(basePath string) *Manager{
	return &Manager{}
}

func (m Manager) Handle() error {
	return nil
}

func isArticleExist(title string) (bool,error){
	return false,nil	
}
