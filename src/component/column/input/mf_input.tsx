import { VFC } from 'react';
import styled from "styled-components";
import {StyledItem, StyledLabel} from 'commonCSS/style';
import { InputProps, InputType } from 'component/column/columnDef'

export const MFInput: VFC<InputProps<InputType>> = ({label, id, input, value, handleChange}) =>{
    return(
        <StyledItem>
            <StyledLabel htmlFor={id}>{label}</StyledLabel>
            <StyledInput
                id={id}
                name={input.name}
                type={input.type}
                defaultValue={value ?? ''}
                required={input.required ?? false}
                placeholder={input.placeholder ?? ''}
                disabled={input.disabled ?? false}
                maxLength={input.maxlength ?? undefined}
                onChange={e => handleChange(e.target.value)}
            >
            </StyledInput>
        </StyledItem>
    )
}

const StyledInput =styled.input`
width: 300px;
height: 20px;
border-radius: 5px;
border: 0.2px solid #9c9c9c;
`;