import { VFC } from 'react';
import styled from "styled-components";
import { TextareaProps } from 'component/column/columnDef'
import {StyledItem, StyledLabel} from 'commonCSS/style';

export const MFTextarea: VFC<TextareaProps> = ({label,id,textarea, value, handleChange}) =>{
    return(
        <StyledItem>
            <StyledLabel htmlFor={id}>{label}</StyledLabel>
                <StyledTextarea
                    id={id}
                    name={textarea.name}
                    defaultValue={value ?? ''}
                    required={textarea.required ?? false}
                    placeholder={textarea.placeolder ?? ''}
                    cols={textarea.cols ?? 20}
                    rows={textarea.rows ?? undefined}
                    maxLength={textarea.maxlength ?? undefined}
                    onChange={e => handleChange(e.target.value)} 
                >
                </StyledTextarea>
        </StyledItem>
    )
}

const StyledTextarea =styled.textarea`
width: 300px;
height: 100px;
border-radius: 5px;
border: 0.2px solid #9c9c9c;
`;