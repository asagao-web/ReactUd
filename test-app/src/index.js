import React, { Component } from "react";
import ReactDOM from "react-dom";
// import AppFooter from "./AppFooter";
import AppFooterFunctionalComponent from "./AppFooterFunctionalComponents";
import AppHeader from "./AppHeader";
import "./index.css";
import AppContent from "./AppContent";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";

class App extends Component {
  constructor(props) {
    super(props);
    this.handlePostChange = this.handlePostChange.bind(this);
    this.state = { posts: [] };
  }

  handlePostChange(posts) {
    this.setState({ posts: posts });
  }

  render() {
    const myProps = {
      title: "My Cool App!",
      subject: "My subject",
      favourite_color: "red",
    };
    return (
      <div className="app">
        <AppHeader
          {...myProps}
          posts={this.state.posts}
          handlePostChange={this.handlePostChange}
        />
        <AppContent
          posts={this.state.posts}
          handlePostChange={this.handlePostChange}
        />
        {/* <AppFooter /> */}
        <AppFooterFunctionalComponent myProperty={"This is Property."} />
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("root"));
