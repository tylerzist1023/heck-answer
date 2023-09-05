import React from 'react';
import {useState, useEffect} from 'react';
import PostComponent from '../components/PostComponent';
import ReplyComponent from '../components/ReplyComponent';

const Post = () => {
    const queryParams = new URLSearchParams(window.location.search);

    const [postid, setPostId] = useState(parseInt(queryParams.get("id")!));
    const [post, setPost] = useState();

    useEffect(() => {
        let res = fetch("/api/post?id="+postid, {
            method: "get"
        })
            .then((r) => r.json())
            .then((json) => {
                setPost(json);
            });
    }, [postid]);
    
    if(post == undefined) {
        return <div></div>;
    } else {
        return (
            <div>
                <PostComponent id={postid} title={post["title"]} author={post["userid"]} body={post["body"]} url={post["url"]}/>
                <ReplyComponent parentid={postid}/>
            </div>
        );
    }
};

export default Post;