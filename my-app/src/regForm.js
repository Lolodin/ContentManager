import React from 'react';
class RegForm extends React.Component {
    render() {return(
        <div className={"reg"}> <h1>Registation Form</h1>
        <form>
<label className={"elem"}>  Login</label><input type={"text"} name={"login"} /> <br/>
<label className={"elem"}>Password</label><input type={"text"} name={"Password"} /><br/>
<label className={"elem"}>Email</label><input type={"text"} name={"Email"} />
        </form>
            <button className={"back btn btn-primary"} onClick={(e)=>this.backMainPage(e)}>Back</button>         <button className={"back btn btn-primary"} onClick={(event)=>this.sendForm(event)} >Registration</button>
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
    backMainPage(event) {
        event.preventDefault()
        let parentState = this.props.getState();
        parentState.page = "mainPage"
        parentState.error= false
        this.props.changeState(parentState)
    }
}
export default RegForm