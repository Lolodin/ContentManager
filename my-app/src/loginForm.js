import React from 'react';

class LoginForm extends React.Component{
    render() {
        return (<div>
            <form action={"localhost:8080/authFunc"} >
                <label>Login</label> <input name={"login"} type={"text"}/>
                <label>Password</label> <input name={"Password"} type={"text"}/>
                <button onClick={()=>this.sendForm()} >Login</button>
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
   async sendForm() {
        let elForm = document.querySelector("form");
        let formdata = new FormData(elForm);
        let response = await fetch("http://localhost:8080/authFunc",{
            method: 'POST',
            body: formdata
        } )
        let js = await response.json()
       console.log(js, "jso455n")
       await this.props.changeState(js)
    }
}
export default LoginForm