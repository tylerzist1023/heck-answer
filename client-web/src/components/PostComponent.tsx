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
        <div className='post'>
            <a target='_blank' href={props.url}>{props.title}</a> by <Link to={"/user?id="+props.author}>{props.author}</Link>

            <br/>
            <Link reloadDocument={true} to={"/post?id="+props.id}>View Thread</Link>
            <div className="display-linebreak">
            {props.body}
            </div>
            <br/>
        </div>
    );
}

export default PostComponent;