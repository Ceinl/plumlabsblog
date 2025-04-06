package articles

type Manager struct {
	basePath string
	article  Article
}

type Article struct {
	Title               string
	ContentMarkdownPath string
	ContentHTMLpath     string
}

func NewArticleManager(basePath string) *Manager{
	return &Manager{
		basePath: basePath,
	}
}

func (m Manager) Handle() error {
	return nil
}

func isDireExist(path string) (bool,error){
	return false,nil	
}

func splitName(filename string) (string, string, error){
	return "","",nil
}
