import { VFC } from 'react';
import styled from "styled-components";
import { Oval } from 'react-loader-spinner'

type ModalSpinnerProps = {
    height: number;
    width: number;
}

export const ModalSppiner: VFC<ModalSpinnerProps> = ({ height, width}) =>{
        return(
            <StyledItem>
                <Oval 
                    color="#F5F5F5"
                    height={height}
                    width={width}
                    secondaryColor='#A9A9A9'
                />
            </StyledItem>
        )
}

const StyledItem =styled.div`
background-color: rgba(0,0,0,0.5); 
width: 100%;
height: 100%;
transition: opacity 0.5s;
align-items: center;
justify-content: center;
display: flex;
position: absolute;
top: 0;
left: 0;
`;