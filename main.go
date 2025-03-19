package main
import (
    "fmt"
)
type WordDictionary struct {
    Alpha map[rune]*WordDictionary
    End bool
    
}

func Constructor() WordDictionary {
    return WordDictionary{Alpha:make(map[rune]*WordDictionary),End:false}
    
}

func (this *WordDictionary) AddWord(word string)  {
    node:=this
    for _,w:=range word{
        _,found:=node.Alpha[w]
        if !found{
            newnode:=Constructor()
            fmt.Println(newnode)
            node.Alpha[w]=&newnode
        }
        node=node.Alpha[w]
    }
    node.End=true
    
}


func (this *WordDictionary) Search(word string) bool {
    node:=this
    for ind,w:=range word{
        if w=='.'{
            for _,val:=range node.Alpha{
                if val.Search(word[ind+1:]){
                    return true
                }
                
            }
            return false

        }else{
            _,found:=node.Alpha[w]
            if !found{
                return false
            }
            node=node.Alpha[w]
        }
    }
    return node.End

}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 **/