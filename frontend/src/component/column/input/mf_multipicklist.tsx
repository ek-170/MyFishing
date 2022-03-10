import { VFC } from 'react';
import styled from "styled-components";
import { MultiPicklistProps } from 'component/column/columnDef'
import {StyledItem, StyledLabel} from 'commonCSS/style';

export const MFMultiPicklist: VFC<MultiPicklistProps> = ({label,id,select,value, options, handleChange}) =>{
    return(
        <StyledItem>
            <StyledLabel htmlFor={id}>{label}</StyledLabel>
            <StyledSelect
                id={id}
                name={select.name}
                multiple
                defaultValue={value}
                required={select.required ?? false}
                onChange={e => handleChange(getValue(e.target.options))}
            >
                {options.map((o)=>(
                    <option value={o.value} key={o.value}>{o.text}</option>
                    ))}
            </StyledSelect>
        </StyledItem>
    )
}

const getValue = (options: HTMLOptionsCollection)=> {
    const values:string[] = [];
    for (let i=0; i<options.length; i++ ){
        if(options[i].selected) values.push(options[i].value);
    }
    return values;
}

const StyledSelect =styled.select`
width: 300px;
height: auto;
border-radius: 5px;
border: 0.2px solid #9c9c9c;
`;