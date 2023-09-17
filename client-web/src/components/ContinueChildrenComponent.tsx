import { useState } from "react";
import ChildrenComponent from "./ChildrenComponent";
import PostComponent from "./PostComponent";


interface ContinueChildrenProps {
    parentid: number;
    indentation: number;
    limit: number;
}

const ContinueChildrenComponent: React.FC<ContinueChildrenProps> = (props: ContinueChildrenProps) => {
    const [continued, setContinued] = useState(false);
    const [children, setChildren] = useState([]);
    const [post, setPost] = useState();

    const continueButtonClicked = (e: any) => {
        e.preventDefault();

        let res = fetch("/api/post/"+props.parentid, {
            method: "get"
        })
            .then((r) => r.json())
                .then((json) => {
                    setPost(json);
                }
        );

        let res2 = fetch("/api/post/"+props.parentid+"/children", {
            method: "get"
        })
            .then((r) => r.json())
                .then((json) => {
                    setChildren(json);
                }
        );
        setContinued(true);
    }

    if(!continued || post == undefined)
    {
        return (
            <button onClick={continueButtonClicked}>Continue This Thread</button>
        );
    }
    else
    {
        return (
            <span>
                <PostComponent
                    id={post["id"]}
                    title={post["title"]}
                    author={post["userid"]}
                    url={post["url"]}
                    body={post["body"]}
                />
                <ChildrenComponent children={children} indentation={props.indentation} parentid={props.parentid} limit={props.limit+10}/>
            </span>
        );
    }
};

export default ContinueChildrenComponent;