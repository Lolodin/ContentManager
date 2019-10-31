import React from 'react';
import ImgCard from "./imgCard";
class MyContent extends React.Component {
    constructor(props) {
        super(props)

    }

   async componentDidMount() {

        let data = {request: "myContent"}
        let req = await fetch("/apiController", {
            method: "POST",
            body: JSON.stringify(data)
        })
         let js = await req.json()
         let images =  Object.entries(js.content)
        this.img =images.map((image, index)=> <ImgCard contentImg = {image[1]} indexIMG={image[0]} key={index} index={index} setDrag={this.setDragEl} setDrop={this.setDropEl}/>)

         this.setState({content:images})
       // this.img =this.state.content.map((image, index)=> <ImgCard contentImg = {image[1]} indexIMG={image[0]} key={index} index={index}/>)

        console.log(this.state)


    }
   render() {

try {
        this.img = this.state.content.map((image, index) => <ImgCard contentImg={image[1]} indexIMG={image[0]}
                                                                     key={index} index={index} setDrag={(e)=>this.setDragEl(e)} setDrop={(e)=>this.setDropEl(e)} leave={()=>this.dnbControllerLeave()} content={this.state.content} deleteImg={(id)=>this.deleteEl(id)}/>)
} catch (e) {
    console.log(e)
}


        return(
            <div>
           <div className={"container"}>
               {this.img}

           </div>
                <button className={"back"} onClick={(event)=>this.backMainPage(event)}>Back</button>
            </div>

        )
    }

setDragEl (el) {
   console.log(el, 'dragstar el');
   this.dragEl = el
}
setDropEl (el) {
    console.log(el);
this.dropEl=el
    this.dnbController(this.dragEl, this.dropEl)
}
dnbController(dragEl, dropEl) {
    let dragE = this.state.content[dragEl]
    let dropE = this.state.content[dropEl]
    this.state.content[dropEl] = dragE
    this.state.content[dragEl] = dropE
    this.setState({})
    console.log("Обновили стейт",this.state.content)
    this.sortArr()
}
dnbControllerLeave() {
    let dragE = this.state.content[this.dropEl]
    let dropE = this.state.content[this.dragEl]
    this.state.content[this.dropEl] = dropE
    this.state.content[this.dragEl]  = dragE
    this.setState({})
}
deleteEl(id) {
    console.log(this.state.content);
    //this.state.content.
    delete this.state.content[id]
    console.log(this.state.content);
    this.setState({})
}
backMainPage(event) {
        event.preventDefault()
        let parentState = this.props.getState();
        parentState.page = "mainPage"
        this.props.changeState(parentState)
    }


async sortArr() {
        let newArr = [];
        let array =Array.from(this.state.content);
        console.log(array)
        array.forEach((v)=>{newArr.push(v[0])})
        newArr.sort((a,b)=>{return a-b})
        array.forEach((v,i, array)=>array[i][0]=newArr[i])
    let updata = {request: "updata", updata: array}
    let req = await fetch("/apiController", {
        method: "POST",
        body: JSON.stringify(updata)
    })
}
}
export default MyContent;