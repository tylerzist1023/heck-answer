import React from 'react';
import {useState, useEffect} from 'react';
import PostComponent from '../components/PostComponent';
import ReplyComponent from '../components/ReplyComponent';
import ChildrenComponent from '../components/ChildrenComponent';

const Post = () => {
    const queryParams = new URLSearchParams(window.location.search);

    const [postid, setPostId] = useState(parseInt(queryParams.get("id")!));
    const [post, setPost] = useState();

    const [children, setChildren] = useState([]);

    const rerender = () => {};

    const refetchChildren = () => {
        let res = fetch("/api/post/"+postid, {
            method: "get"
        })
            .then((r) => r.json())
                .then((json) => {
                    setPost(json);
                }
        );

        let res2 = fetch("/api/post/"+postid+"/children", {
            method: "get"
        })
            .then((r) => r.json())
                .then((json) => {
                    setChildren(json);
                }
        );
    };
    
    useEffect(refetchChildren, [postid])

    if(post == undefined) {
        return <div></div>;
    } else {
        return (
            <div>
                <PostComponent id={postid} title={post["title"]} author={post["userid"]} body={post["body"]} url={post["url"]}/>
                <ReplyComponent parentid={postid} refetchChildren={refetchChildren}/>

                <ChildrenComponent parentid={postid} children={children} indentation={0} limit={10} />
            </div>
        );
    }
};

export default Post;