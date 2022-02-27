import { VFC } from 'react';
import { ButtonProps,ButtonType, ButtonMode } from 'component/button/buttonDef'
import styled from "styled-components";
import character from 'const/character';

export const MFButton: VFC<ButtonProps<ButtonType, ButtonMode>> = ({type, mode, children, disabled ,onClick, hidden}) =>{
    return(
        <StyledButton type={type} disabled={disabled} onClick={onClick} mode={mode} hidden={hidden}>
            {children}
        </StyledButton>
    )
}

const StyledButton =styled.button.attrs<ButtonProps<ButtonType, ButtonMode>>((props)=> ({
    mode: props.mode
}))<ButtonProps<ButtonType, ButtonMode>>`
    display: inline-block;
    padding: 5px 17px;
    background-color: ${props => props.mode==='positive' ? '#6495ED' : '#808080'};
    color:  white;
    font-family: ${character.FONT};
    font-size:  ${character.SIZE.LARGE}px;
    border-radius: 3px;
    border-style: none;
    cursor:pointer;
    transition: 0.2s ;
    &:hover{
        background-color: ${props => props.mode==='positive' ? '#87CEFA' : '#D3D3D3'};
    `;