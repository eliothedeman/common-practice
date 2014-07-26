package search

import (
	"testing"
)

const (
	TEST_KEY             = "test"
	TEST_DATA            = "laksjdflaksjdflaskjdfl;aksjaksjdfl;akjsdlfkjkajiwuerinwenenwehuuxZJclnwehrherjnwfenaldjfajsdhfkjjjdafsdjfashdfahsdiufhaiwuenncaiwdsuhfiahweurinkjnwdfkjansdufhau9wenfuansdkjfnsdkbaksjdghkasjdfl;ads;flahsdifuhasdfnawdjnwkejnfiawdbgkjasdfkjahsdlfjasldkfjasldghauisdhfunfadsjfnakljbwerkjbakdjchajsdhfkjabdfkjlabdskhfhaksjdfhakjsdhfkjashdfjahcjdncandiufahsdkljf asdjfha dfa sdfasdf asdfhasdf ashdfjasdhfshdf asdfhajsdhijfhaklwjehkcn asd adfhajsdf adsg "
	SELECT_ALL_TEST_DATA = "jkjk kkk jkjk kkj"
)

var (
	s *SearchDb
)

// TestSplitData tests the TestSplitData function
func TestSplitData(t *testing.T) {
	d := "asdfgasdfasdkfjasldkjfal;sdjf;alsdkf"
	SplitData(d)
}

// TestInsert tests the Insert function
func TestInsert(t *testing.T) {
	s = NewSearchDb("testing")
	s.Insert("test", 1)
}

// TestSelect tests the Select function
func TestSelect(t *testing.T) {
	expected := 1
	if expected != s.Select("test") {
		t.Fail()

	}
}

// TestSelectAll tests selecting all of the KeyPointers of a given key
func TestSelectAll(t *testing.T) {
	one := NewSearchDb("one")
	err := one.CreateAndUpdateindicies(SELECT_ALL_TEST_DATA)
	if err != nil {
		t.Error(err)
	}
	keys := one.SelectAll("j")
	println(len(keys))
}

// TestMultipleManagers makes sure multiple managers don't have cross-talk
func TestMultipleManagers(t *testing.T) {
	one := NewSearchDb("one")
	two := NewSearchDb("two")
	one.Insert("test", "one")
	two.Insert("test", "two")
	if one.Select("test") != "one" {
		t.Fail()
	}
	if two.Select("test") != "two" {
		t.Fail()
	}
}

// TestUpdateChildren makes sure we can still update the children of an index
func TestUpdateChildren(t *testing.T) {
	// create the database for testing
	one := NewSearchDb("one")
	// insert a new index
	one.Insert(TEST_KEY, NewIndex())
	// if we can't update, error out
	err := one.UpdateChildren(TEST_KEY, &KeyPointer{Key: TEST_KEY, Pointer: 1})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// make sure the data got added
	index := one.Select(TEST_KEY)
	// check that it is the correct type, and it has the correct value
	if len(index.(*Index).Children) != 1 {
		t.Log("Child not added. Wrong length: ", len(index.(*Index).Children))
		t.Fail()
	}
}

// TestCreateIndexIfNotExist tests that indicies are not overwritten by this method
func TestCreateIndexIfNotExist(t *testing.T) {
	one := NewSearchDb("one")
	one.CreateIndexIfNotExist("1")
	one.UpdateChildren("1", &KeyPointer{Key: TEST_KEY, Pointer: 1})
	one.CreateIndexIfNotExist("1")
	if len(one.Select("1").(*Index).Children) != 1 {
		t.Log("Overwrote Index.")
		t.Fail()
	}
}

// TestCreateAndUpdateindicies makes sure a dataset can be created in the db
func TestCreateAndUpdateindicies(t *testing.T) {
	one := NewSearchDb("one")
	one.CreateAndUpdateindicies("laksjdfklajsdflkajsd;lfjka")
}

// BenchmarkTestCreateAndUpdateindicies
func BenchmarkTestCreateAndUpdateindicies(b *testing.B) {
	one := NewSearchDb("one")
	for i := 0; i < b.N; i++ {

		one.CreateAndUpdateindicies(TEST_DATA)
	}

}

// BenchmarkSplitData tests performance of SplitData
func BenchmarkSplitData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitData(TEST_DATA)
	}
}