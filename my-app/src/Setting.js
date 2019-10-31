import React from 'react';

class Setting extends React.Component{
    constructor() {
        super()
        this.state = {elementValue: []}
    }
async componentDidMount() {
        console.log("getSetting")
  await  this.getUserSetting()
}

    render() {

        return(

            <div className={"login"}>
                <form>
                    <label className={"elem"}>Chat id:</label> <input type={"text"} name={"chatId"} defaultValue={this.state.elementValue.chatid} />
                    <br/>
                    <label className={"elem"}>Hours</label> <input type={"number"} name={"hours"}   defaultValue={this.state.elementValue.hours}/>
                    <br/>
                    <label className={"elem"}>Minutes</label> <input type={"number"} name={"minutes"} defaultValue={this.state.elementValue.minutes}/>
                </form>
                <button className={"back btn btn-primary"}  onClick={(e)=>this.backMainPage(e)}>Back</button> <button className={"back btn btn-primary"}  onClick={()=>this.sendData()}>Send</button>
            </div>
        )
    }
    sendData() {

        let form = document.querySelector("form")
        let fd = new FormData(form)
        let array = {}
        fd.forEach((v,k)=>{
            console.log(k,v)
            array[k] = v
        })
        console.log(array)
        let data = {request: "setting", contentx: array}
        let req =  fetch("/apiController", {
            method: "POST",
            body: JSON.stringify(data)
        })



    }
    backMainPage(event) {
        event.preventDefault()
        let parentState = this.props.getState();
        parentState.page = "mainPage"
        this.props.changeState(parentState)
    }
  async  getUserSetting() {
        let request = {request: "getSetting"}
        let req = await fetch("/apiController", {
            method: "POST",
            body: JSON.stringify(request)
        } )
        let res = await req.json()

this.setState({elementValue: res.setting})
console.log(this.state)

    }
}
export default Setting