package nihongo

import (
	"fmt"

	sql "github.com/FloatTech/sqlite"
)

type grammar struct {
	ID            int    `db:"id"`
	Tag           string `db:"tag"`
	Name          string `db:"name"`
	Pronunciation string `db:"pronunciation"`
	Usage         string `db:"usage"`
	Meaning       string `db:"meaning"`
	Explanation   string `db:"explanation"`
	Example       string `db:"example"`
	GrammarURL    string `db:"grammar_url"`
}

func (g *grammar) string() string {
	return fmt.Sprintf("ID:\n%d\n\n标签:\n%s\n\n语法名:\n%s\n\n发音:\n%s\n\n用法:\n%s\n\n意思:\n%s\n\n解说:\n%s\n\n示例:\n%s", g.ID, g.Tag, g.Name, g.Pronunciation, g.Usage, g.Meaning, g.Explanation, g.Example)
}

var db = &sql.Sqlite{}

func getRandomGrammarByTag(tag string) (g grammar) {
	_ = db.Find("grammar", &g, "where tag LIKE '%"+tag+"%' ORDER BY RANDOM() limit 1")
	return
}

func getRandomGrammar() (g grammar) {
	_ = db.Pick("grammar", &g)
	return
}
