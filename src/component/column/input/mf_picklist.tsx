import { VFC } from 'react';
import styled from "styled-components";
import { PicklistProps } from 'component/column/columnDef'
import {StyledItem, StyledLabel} from 'commonCSS/style';

export const MFPicklist: VFC<PicklistProps> = ({label,id,select,value, options, handleChange}) =>{
    return(
        <StyledItem>
            <StyledLabel htmlFor={id}>{label}</StyledLabel>
            <StyledSelect
                id={id}
                name={select.name}
                defaultValue={value ?? undefined}
                required={select.required ?? false}
                onChange={e => handleChange(e.target.value)}
            >
                <option hidden>選択してください</option>
                {options.map((o)=>(
                    <option value={o.value} key={o.value}>{o.text}</option>
                    ))}
            </StyledSelect>
        </StyledItem>
    )
}

const StyledSelect =styled.select`
width: 300px;
height: auto;
border-radius: 5px;
border: 0.2px solid #9c9c9c;
`;