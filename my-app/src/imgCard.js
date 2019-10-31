import React from 'react';
class ImgCard extends React.Component{
constructor(props) {
    super(props)
this.state ={styleImg: {order: this.props.index, opacity: 1}}
}


render() {


    return(
        <div onDragOver={()=>this.dragOver()} onDragLeave={()=>this.dragLeave()} className={"item"} draggable={true} style={this.state.styleImg} id={this.props.indexIMG}>
            <button className={'delete'} onClick={()=>this.deletBD()}>X</button>
            <img  onDragEnter={(event)=>this.dragEnter(event)}  onDragStart={(event)=>this.dragStart(event)} onDragEnd={(event)=>this.dragEnd(event)} onDrop={(event)=>this.dragEnd(event)}  src={"/image/" + this.props.contentImg} height={128} width={128} />
            <p>{this.props.index}</p>
        </div>
    )

}
dragStart(ev) {
    this.props.setDrag(this.props.index)
    this.setState( {styleImg: {order: this.props.index, opacity:0.30}})
}
dragEnd(e) {
    this.props.setDrop( this.props.index)
}
dragEnter(e) {
    this.props.setDrop( this.props.index)
     console.log('над элементом',  this.props.index)
}
dragLeave() {
    this.setState( {styleImg: {order: this.props.index, opacity: 1}})
}
dragOver() {
    this.setState( {styleImg: {order: this.props.index, border: '3px solid orange'}})
}

async deletBD() {
    console.log(this.props.content[this.props.index][0], 'delete index')
    let data = {request: "delete", content: this.props.content[this.props.index]}
    let req = await fetch("/apiController", {
        method: "POST",
        body: JSON.stringify(data)
    })
this.props.deleteImg(this.props.index)

}
}
export default ImgCard;