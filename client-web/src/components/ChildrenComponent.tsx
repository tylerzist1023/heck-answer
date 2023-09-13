import React, { useState } from 'react';
import {
  Link
} from "react-router-dom";
import PostComponent from './PostComponent';

interface ChildrenProps {
    parentid: number;
    children: any;
    indentation: number;
}

interface HiddenType {
    [id: string]: boolean;
}

const ChildrenComponent: React.FC<ChildrenProps> = (props: ChildrenProps) => {
    const [hidden, setHidden] = useState({} as HiddenType);

    const toggleHidden = (e: any, id: string) => {
        e.preventDefault();

        if(hidden[id] == undefined)
        {
            var newHidden = {...hidden, [id]: true};
            setHidden(newHidden);
        }
        else
        {
            var newHidden = {...hidden};
            if(hidden[id] == false)
            {
                newHidden[id] = true;
            }
            else
            {
                newHidden[id] = false;
            }
            setHidden(newHidden);
        }
    }

    return (
        <div>
            {
                props.children.filter((e: any,i: number,a: any) => { return e["parentid"]==props.parentid; }).map((x: any) => {

                    var grandchildren = props.children.filter((e:any,i:number,a:any) => {return e["parentid"]==x["id"] })

                    return (
                        <div>
                            <div className='post-child' style={{marginLeft: 30+'px'} }>
                                <button onClick={(e: any) => toggleHidden(e, x["id"] as string)}>Hide or Show</button>
                            {
                                hidden[x["id"] as string]==true ? <span></span> : 
                                    <div>
                                        <PostComponent
                                            id={x["id"]}
                                            title={x["title"]}
                                            author={x["userid"]}
                                            url={x["url"]}
                                            body={x["body"]}
                                        />
                                        {
                                            grandchildren.length > 0 ? 
                                                <ChildrenComponent  parentid={x["id"]} children={props.children} indentation={props.indentation+1}/> : 
                                                <span></span>
                                        }
                                    </div>
                            }
                            </div>
                            { props.indentation==0 ? <hr/> : <span></span> }
                        </div>
                    );
                })
            }
        </div>
    );
}

export default ChildrenComponent;