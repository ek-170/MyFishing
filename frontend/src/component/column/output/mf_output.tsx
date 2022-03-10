import { VFC } from 'react';
import styled from "styled-components";
import { BaseColumnProps } from 'component/column/columnDef'
import {StyledItem, StyledLabel} from 'commonCSS/style';
import { ValueType } from 'component/column/columnDef'

export const MFOutoPut: VFC<BaseColumnProps> = ({label, id, value}) =>{
    value = isMultiPickList(value);
    return(
        <StyledItem>
            <StyledLabel id={id}>{label}</StyledLabel>
                <StyledValue> { value }</StyledValue>
        </StyledItem>
    )
}

const isMultiPickList = (value:ValueType | undefined):ValueType | undefined => {
    if( !Array.isArray(value)){ return value; }

    let result:string = '';
    for(let v of value){
        result = result + v + ',';
    }
    return result.slice(0, -1);
}

const StyledValue =styled.div`
margin-top: 5px;
width: 300px;
`;