import React from 'react';
import logo from './logo.svg';
import RegForm from "./regForm.js";
import LoginForm from "./loginForm.js";
import './App.css';
import SendFile from "./sendFile.js";


class mainPage extends React.Component {
    constructor(props) {
        super(props)
       this.state = {page: "mainPage", login: "false"} //test login

    }

    render() {
        if (this.state.page == "mainPage" && this.state.login != "false") {
            return (<div><h1>Work Page</h1>
                    <h2>Страница тестирования Json запросов</h2>
                    <h2 onClick={()=>this.actionLoadFile()}>Добавить контент</h2>
                    <h2>Настройки</h2>
                    <h2 onClick={()=>this.exitSsesion()}>Выход</h2>
                </div>
            )
        }
        if (this.state.page == "mainPage" && this.state.login == "false") {
            return (<div><h1>Main Page</h1>
                    <h2 onClick={()=>this.actionLoginPage()}>Login</h2>
                    <h2 onClick={()=>this.actionRegPage()}>Registration</h2>
                </div>
            )
        }
        if (this.state.page == "regPage" && this.state.login == "false") {
            return (<RegForm />)
        }

        if (this.state.page == "authPage" && this.state.login == "false") {
            return (<LoginForm changeState={this.changeState} />)
        }
        if (this.state.page == "sendFile" && this.state.login != "false") {
            return (<SendFile changeState = {this.changeState} getState ={this.getState}/>)
        }
    }
   async componentDidMount()
    {
       // fetch('http://localhost:8080/checkAuth').then((response)=> response.json().then((response)=> this.setState(response)))
        let response = await fetch('http://localhost:8080/checkAuth');
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
     console.log(this.state, "old")
     console.log(newState)
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
       document.cookie = "user = ;max-age=0"
        this.setState({page:"mainPage", login: "false"})

    }
    actionLoadFile() {
        this.setState({page:"sendFile"})
    }
}

export default mainPage;
