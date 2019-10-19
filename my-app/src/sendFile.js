import React from 'react';
class SendFile extends React.Component {
    render() {
        return (<form action={"http://localhost:8080/formHandler"} method={"POST"} encType={"multipart/form-data"}>
            <label >Text</label>  <input name={"testText"} type={"text"} />
            <br/>
            <label>File: </label><input name={"testFile"} type={"file"}/>
            <button onClick={(event)=>this.sendfile(event)}>Send File</button>
            <button onClick={(event)=>this.backMainPage(event)}>Back</button>

        </form>)
    }
    backMainPage(event) {
        event.preventDefault()
       let parentState = this.props.getState();
        parentState.page = "mainPage"
        this.props.changeState(parentState)
    }
   async sendfile(event) {
       event.preventDefault()
        let form = document.querySelector("form")
        let fd = new FormData(form)
        let request = await fetch("/formHandler",{
            method: "POST",
            body: fd,
        })
       let js = await request.json();
       this.props.changeState(js);


    }
}
export default SendFile