import React from 'react';
import {useState, useEffect} from 'react';
import { useNavigate } from 'react-router';
import PostComponent from './PostComponent';

interface ReplyProps {
    parentid: number;
    refetchChildren: () => void;
}

const ReplyComponent: React.FC<ReplyProps> = (props: ReplyProps) => {
    const navigate = useNavigate();

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        let formData = new FormData(e.currentTarget);
        formData.append("parentid", props.parentid.toString())

        let res = fetch("/api/post/", {
            body: formData,
            method: "post"
        }).finally(() => {
            props.refetchChildren();
        });
    };

    return (
        <div>
            <form method="post" onSubmit={onSubmit}>
                <br />

                <label>
                Reply:
                <textarea id="body" name="body" rows={10} cols={100}></textarea>
                </label>

                <br />

                <input type="submit" value="Reply!" />
            </form>
        </div>
    );
}

export default ReplyComponent;