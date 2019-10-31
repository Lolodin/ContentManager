import React from 'react';
import RegForm from "./regForm.js";
import LoginForm from "./loginForm.js";
import './App.css';
import SendFile from "./sendFile.js";
import MyContent from "./MyContent";
import Setting from "./Setting";


class mainPage extends React.Component {
    constructor(props) {
        super(props)
       this.state = {page: "mainPage", login: "false"} //test login

    }

    render() {
        if (this.state.page == "mainPage" && this.state.login != "false") {
            return (<div className={"mainpage"}><h1 className={"headlogin"}>Work Page</h1>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=>this.actionLoadFile()}>Добавить контент</button>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=> this.changeState({page: "myContent", error: false}) } changeState = {this.changeState}  getState ={this.getState} >Мой контент</button>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=> this.changeState({page: "setting", error: false})}>Настройки</button>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=>this.exitSsesion()}>Выход</button>               </div>
            )
        }
        if (this.state.page == "mainPage" && this.state.login == "false") {
            return (<div className={"login"}><h1 className={"headlogin"}>Main Page</h1>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=>this.actionLoginPage()}>Login</button>
                    <br/>
                    <button type="button" class="btn mainmenu btn-primary" onClick={()=>this.actionRegPage()}>Registration</button>
                </div>
            )
        }
        if (this.state.page == "regPage" && this.state.login == "false") {
            return (<RegForm  changeState = {this.changeState} getState ={this.getState} />)
        }

        if (this.state.page == "authPage" && this.state.login == "false") {
            return (<LoginForm  changeState = {this.changeState} getState ={this.getState} />)
        }
        if (this.state.page == "sendFile" && this.state.login != "false") {
            return (<SendFile changeState = {this.changeState} getState ={this.getState}/>)
        }
        if (this.state.page == "myContent" && this.state.login != "false") {
            return (<MyContent changeState = {this.changeState} getState ={this.getState}/>)
        }
        if (this.state.page == "setting" && this.state.login != "false") {
            return (<Setting changeState = {this.changeState} getState ={this.getState}/>)
        }
    }
   async componentDidMount()
    {
       // fetch('http://localhost:8080/checkAuth').then((response)=> response.json().then((response)=> this.setState(response)))
        let response = await fetch('/checkAuth');
        let json = await response.json();
        this.setState(json);
    }
    // getAuth() {
    //
    //         let req = new XMLHttpRequest();
    //         let state
    //         req.open("POST", "http://localhost:8080/checkAuth")
    //         req.onload= ()=>{
    //
    //             if(req.status==200) {
    //               state = JSON.parse(req.responseText)
    //                 console.log(state)
    //                 return state
    //             }
    //             else
    //             {
    //               state = {page: "mainPage", login: false}
    //                 return state
    //             }
    //         }
    //         req.send()
    //
    //
    //
    // }
 changeState = (newState) => {
 if (newState.error != false) {
     alert("Error: "+newState.error)
     return
 }
        this.setState(newState)


    }
 getState = ()=> {
        return this.state
    }
    actionLoginPage() {
        this.setState({page: "authPage"})
    }

    actionRegPage() {
        this.setState({page: "regPage"})


    }
    exitSsesion() {
        let e = window.confirm('Выполнить выход?');
        if (!e) {
            return
        }
       document.cookie = "user = ;max-age=0"
        this.setState({page:"mainPage", login: "false"})

    }
    actionLoadFile() {
        this.setState({page:"sendFile"})
    }
}

export default mainPage;
