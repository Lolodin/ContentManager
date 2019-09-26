import React from 'react';
class RegForm extends React.Component {
    render() {return(
        <div> <h1>Registation Form</h1>
        <form method={"POST"} action={"localhost:8080/regfunc"}>
<label>Login</label><input type={"text"} name={"login"} /> <br/>
<label>Password</label><input type={"text"} name={"Password"} /><br/>
<label>Email</label><input type={"text"} name={"Email"} />
            <button onClick={()=>this.sendForm()} >Registration</button>
        </form>
        </div>)
    }

    sendForm() {
        let elForm = document.querySelector("form");
        let formdata = new FormData(elForm);
        let req = new XMLHttpRequest();
        req.open("POST", "http://localhost:8080/regfunc")
        req.onload= ()=>{
            console.log(req.status)
            if(req.status==200) {
                alert(req.responseText)
            }
            else
            {
                alert("error")
            }
        }
        req.send(formdata)

    }
}
export default RegForm