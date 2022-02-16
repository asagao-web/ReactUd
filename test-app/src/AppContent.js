import React, { Component } from "react";

export default class AppContent extends Component {
  constructor(props) {
    super(props);
    this.handlePostChange = this.handlePostChange.bind(this);
  }

  handlePostChange(posts) {
    this.props.handlePostChange(posts);
  }

  clickedItem = (x) => {
    console.log("clicked", x);
  };

//   state = { posts: [] };

  fetchList = () => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((json) => {
        // this.setState({ posts: json });
        this.handlePostChange(json);
      });
  };
  render() {
    return (
      <div>
        This is the AppContent.
        <br />
        <hr />
        <button onClick={this.fetchList} className="btn btn-primary" href="#">
          Fetch Data
        </button>
        <hr />
        <p>Post is {this.props.posts.length} items long</p>
        <ul>
          {this.props.posts.map((c) => (
            <li key={c.id}>
              <a href="#!" onClick={() => this.clickedItem(c.id)}>
                {c.title}
              </a>
              <p onClick={() => this.clickedItem(c.id)}>
                {" "}
                also can click here.
              </p>
            </li>
          ))}
        </ul>
      </div>
    );
  }
}
