import React from 'react';

class MyContent extends React.Component {
    constructor(props) {
        super(props)

    }

    async componentWillMount() {

        let data = {request: "myContent"}
        let req = await fetch("/apiController", {
            method: "POST",
            body: JSON.stringify(data)
        })
         let js = await req.json()
        let images = await js.content;
        this.images = await images.map((image)=> <img src={"/image/" + image} alt={"MyImage"}/>);
        console.log(images)
        console.log(this.images)
        this.setState({})

    }
   render() {
        return(
           <div>
               {this.images}
           </div>
        )
    }


}
export default MyContent;