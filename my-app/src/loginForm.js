import React from 'react';

class LoginForm extends React.Component{
    render() {
        return (<div className={"auth"}>
            <form action={"/authFunc"} >
                <label className={"elem"}>Login</label> <input name={"login"} type={"text"}/>
                <br/>
                <label className={"elem"}>Password</label> <input name={"Password"} type={"text"}/>
                <br/>
                <br/>
                <button className={"back btn btn-primary"} onClick={(e)=>this.backMainPage(e)}>Back</button>    <button className={"btn btn-primary"} onClick={(event)=>this.sendForm(event)} >Login</button>
            </form>
        </div>)
    }

   async sendForm(event) {
       event.preventDefault()
        let elForm = document.querySelector("form");
        let formdata = new FormData(elForm);
        let response = await fetch("/authFunc",{
            method: 'POST',
            body: formdata
        } )
        let js = await response.json()
       console.log(js, "jso455n")
       await this.props.changeState(js)
    }
    backMainPage(event) {
        event.preventDefault()
        let parentState = this.props.getState();
        parentState.page = "mainPage"
        this.props.changeState(parentState)
    }
}
export default LoginForm