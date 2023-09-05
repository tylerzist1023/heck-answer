import React from 'react';
import {useState, useEffect} from 'react';
import { useNavigate } from 'react-router';
import PostComponent from './PostComponent';

interface ReplyProps {
    parentid: number;
}

const ReplyComponent: React.FC<ReplyProps> = (props: ReplyProps) => {
    const navigate = useNavigate();

    const [children, setChildren] = useState([]);

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        let formData = new FormData(e.currentTarget);
        formData.append("parentid", props.parentid.toString())

        let res = fetch("/api/post/", {
            body: formData,
            method: "post"
        }).finally(() => {
            window.location.reload();
        });
    };

    useEffect(() => {
        let res = fetch("/api/children?id="+props.parentid, {
            method: "get"
        })
        .then((r) => r.json())
        .then((json) => {
            setChildren(json);
        });
    }, []);

    return (
        <div>
            <form method="post" onSubmit={onSubmit}>
                <br />

                <label>
                Reply:
                <textarea id="body" name="body"></textarea>
                </label>

                <br />

                <input type="submit" value="Reply!" />
            </form>

            {
                children.map((x) => { return (<PostComponent 
                    id={x["id"]}
                    title={x["title"]}
                    author={x["userid"]}
                    url={x["url"]}
                    body={x["body"]}
                />);})
            }
        </div>
    );
}

export default ReplyComponent;