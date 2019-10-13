import React from 'react';
class RegForm extends React.Component {
    render() {return(
        <div> <h1>Registation Form</h1>
        <form>
<label>Login</label><input type={"text"} name={"login"} /> <br/>
<label>Password</label><input type={"text"} name={"Password"} /><br/>
<label>Email</label><input type={"text"} name={"Email"} />
            <button onClick={(event)=>this.sendForm(event)} >Registration</button>
        </form>
        </div>)
    }
// переделать под fetch
  async  sendForm(event) {
        event.preventDefault()
        let elForm = document.querySelector("form")
        let formdata = new FormData(elForm)
        let request  = await fetch("/regfunc", {
        method: "POST",
        body: formdata,
        })
      let js = await request.json()
      console.log(js)
      await this.props.changeState(js)
  }
}
export default RegForm