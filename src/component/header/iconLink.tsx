import { VFC } from 'react';
import styled from "styled-components";
import { Link } from "react-router-dom";

type IconLinkProps = {
    to: string,
    src: string,
    alt?: string,
    width?: string,
}

export const IconLink:VFC<IconLinkProps> = ({to,src,alt,width})=> {
    return(
        <Link to={to}><StyledIconLink src={src} alt={alt} width={width}></StyledIconLink></Link>
    )
}

const StyledIconLink =styled.img`
    width:${(props)=>props.width};
    height:auto;
    margin-right:7px;
    padding: 8px 16px;
&:hover{
    background: #383d43;
	transition: 0.2s ;
}
`