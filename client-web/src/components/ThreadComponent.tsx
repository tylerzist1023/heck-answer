import React from 'react';
import {
  Link
} from "react-router-dom";

interface ThreadProps {
    id: number;
    title: string;
    author: string;
    url: string;
}

const ThreadComponent: React.FC<ThreadProps> = (props: ThreadProps) => {
    return (
        <p>
        <a target='_blank' href={props.url}>{props.title}</a> by <Link to={"/user?id="+props.author}>{props.author}</Link>
        <br/>
        <Link to={"/post?id="+props.id}>View Thread</Link>
        </p>
    );
}

export default ThreadComponent;