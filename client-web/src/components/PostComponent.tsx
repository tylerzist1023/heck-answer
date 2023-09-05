import React from 'react';
import {
  Link
} from "react-router-dom";

interface PostProps {
    id: number;
    title: string;
    author: string;
    url: string;
    body: string;
}

const PostComponent: React.FC<PostProps> = (props: PostProps) => {
    return (
        <p>
        <a target='_blank' href={props.url}>{props.title}</a> by {props.author}
        <br/>
        <p>
        {props.body}
        </p>
        </p>
    );
}

export default PostComponent;