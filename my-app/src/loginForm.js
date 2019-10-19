import React from 'react';

class LoginForm extends React.Component{
    render() {
        return (<div>
            <form action={"/authFunc"} >
                <label>Login</label> <input name={"login"} type={"text"}/>
                <label>Password</label> <input name={"Password"} type={"text"}/>
                <button onClick={(event)=>this.sendForm(event)} >Login</button>
            </form>
        </div>)
    }
    // sendForm() {
    //     let elForm = document.querySelector("form");
    //     let formdata = new FormData(elForm);
    //     let req = new XMLHttpRequest();
    //     req.open("POST", "http://localhost:8080/authFunc")
    //     req.onload= ()=>{
    //         console.log(req.status)
    //         if(req.status==200) {
    //             alert(req.responseText)
    //         }
    //         else
    //         {
    //             alert("error")
    //         }
    //     }
    //     req.send(formdata)
    //
    // }
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
}
export default LoginForm