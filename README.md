# Log in your own style
*   My [LinkedIn](https://www.linkedin.com/in/cigeralex/)
*   My wechat account: neverexistsnever
*   My QQ account: 310276663

##### I used to chose golang/glog to log things.It's a very nice package, but that's not what I want. So I try to develop a tool to log messages in a way I want.
---
## Flags

Use **-info** to turn on _log.Info*()_ prints

Use **-warn** to turn on _log.Warning*()_ prints

Use **-err false** to turn off _log.Error*()_ prints

#### In developing:

##### Use **-time** to print message with a _time.Now().UTC().Format(time.RFC3339)_ tag.

_log.Error*()_ and _log.Error*()_ prints are on by default.

---
## Print in info, warning and error styles:

* Use _Infof()_, _Warningf()_, _Errorf()_ to print informations **INLINE**, and use them just like you use _fmt.Printf()_.

* Example[1]:

        1  package main
        2  import "github.com/achillesss/log"
        3  func print(){
        4    log.Parse()
        5    log.Infof("My name is %s.", Alex) // simply an information
        6    log.Warningf("Tell me how to spell %q, please.", golang) // pretend to be warned
        7    log.Errorf("g-o-l-a-n.") // that's exactly an error
        8  }
        9  func main() {
        10     log.Parse()
        11     print()
        12 }
and you type 'go run main.go -info, -warn' in cmd-line
then you'll get

        [I_main.go_5] My name is Alex.    [W_main.go_6] Tell me how to spell "golang", please.    [E_main.go_7] g-o-l-a-n.
all messages above inline.

But why do this?

Just try tail -f *.log | grep xxx | grep xxx
* Use _Infofln(), _Warningfln()_, _Errorfln()_ to print informations like _fmt.Printf(format+"\n",args...)_

Replace each _log.Infof_, _log.Warningf_, _log.Errorf_ by _log.Infofln_, _log.Warninfln_, _log.Errorfln_, and you'll get

        [I_main.go_5] My name is Alex.
        [W_main.go_6] Tell me how to spell "golang", please.
        [E_main.go_7] g-o-l-a-n.
But why do this?

Want you to debug with all messages **inline**? I don't know and I don't want.

---
## Log:

Time flies. I'll add this field later.

---
## Error Formating:

Time flies. I'll add this field later.

---
## Print name of your _func()_:

Time flies. I'll add this field later.

---
## Versions:
* #### 1.0:
    Add basic funtions
* #### 2.0:
    Rebuild and add some toggles
* #### 2.0.1:
    Add README.md