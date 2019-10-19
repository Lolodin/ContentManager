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
                                                                     key={index} index={index} setDrag={(e)=>this.setDragEl(e)} setDrop={(e)=>this.setDropEl(e)} />)
} catch (e) {
    console.log(e)
}


        return(
           <div className={"container"}>
               {this.img}
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
   // console.log(this.state.content[dragEl], 'перемещаемый')
    let dragE = this.state.content[dragEl]
   // console.log(this.state.content[dropEl], 'принимаемый')
    let dropE = this.state.content[dropEl]
    this.state.content[dropEl] = dragE
    this.state.content[dragEl] = dropE
    this.setState({})
    console.log(this.state)
}

}
export default MyContent;