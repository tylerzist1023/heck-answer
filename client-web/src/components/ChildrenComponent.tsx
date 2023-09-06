import React from 'react';
import {
  Link
} from "react-router-dom";
import PostComponent from './PostComponent';

interface ChildrenProps {
    parentid: number;
    children: any;
    indentation: number;
}

const ChildrenComponent: React.FC<ChildrenProps> = (props: ChildrenProps) => {
    return (
        <div>
            {
                props.children.filter((e: any,i: number,a: any) => { return e["parentid"]==props.parentid; }).map((x: any) => {

                    var grandchildren = props.children.filter((e:any,i:number,a:any) => {return e["parentid"]==x["id"] })

                    return (
                        <div style={{marginLeft: 30+'px'} }>
                            <PostComponent
                                id={x["id"]}
                                title={x["title"]}
                                author={x["userid"]}
                                url={x["url"]}
                                body={x["body"]}
                            />
                            {
                                grandchildren.length > 0 ? 
                                    <ChildrenComponent parentid={x["id"]} children={props.children} indentation={props.indentation+1}/> : 
                                    <br/>
                            }
                        </div>
                    );
                })
            }
        </div>
    );
}

export default ChildrenComponent;