import styled from "styled-components";
import { header, container } from 'const/size';
import character from 'const/character';
import { color } from 'const/color';

/* 
ページ全体
 */
export const StyledOuter =styled.div`
    background-color: #eff0f0;
    width: 100vw;
    min-height: 100vh;
`;

export const StyledWrapper =styled.div`
    margin-top: ${header.HEIGHT};
    width: ${container.WIDTH};
    min-height: 100vh;
    margin-left: auto;
    margin-right: auto;
    background-color: #ffffff;
`;

/* 
項目関連
 */
export const StyledItem =styled.div`
margin-left: auto;
margin-right: auto;
font-family: ${character.FONT};
font-size:  ${character.SIZE.LARGE}px;
`;

export const StyledLabel =styled.label`
display: inline-block;
vertical-align: top;
display: flex;
align-items: center;
height: 40px
border: solid 5px #5250e0;
color: ${color.COLUMN.LABEL};
`;