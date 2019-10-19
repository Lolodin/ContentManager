import React from 'react';
class ImgCard extends React.Component{
constructor(props) {
    super(props)
this.state ={styleImg: {order: this.props.index, opacity: 1}}
}


render() {


    return(
        <div className={"item"} draggable={true} style={this.state.styleImg} id={this.props.indexIMG}>
            <img onDragEnter={(event)=>this.dragEnter(event)}  onDragStart={(event)=>this.dragStart(event)} onDragEnd={(event)=>this.dragEnd(event)}   src={"/image/" + this.props.contentImg} height={128} width={128} />
            <p>{this.props.index}</p>
        </div>
    )

}
dragStart(ev) {

    this.props.setDrag(this.props.index)

    this.setState( {styleImg: {order: this.props.index, opacity:0.30}})

}
dragEnd(e) {

    this.setState( {styleImg: {order: this.props.index, opacity: 1}})
}
dragEnter(e) {

    this.props.setDrop( this.props.index)
    console.log('над элементом',  this.props.index)
}
}
export default ImgCard;