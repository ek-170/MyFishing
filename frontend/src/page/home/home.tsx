import { VFC } from 'react';
import { Helmet, HelmetProvider } from 'react-helmet-async';
import HomeImage from 'img/HomeBackgroundImage.jpg';
import styled from "styled-components";

const Home: VFC = () => {
    return (
        <HelmetProvider>
            <Helmet>
            <title>Home</title>
            </Helmet>
                <StyledBackgroundImg>
                    <StyledText>My Fishing</StyledText>
                </StyledBackgroundImg>
        </HelmetProvider>
    )
}

const StyledBackgroundImg = styled.div`
position: relative;
width: 100%;
min-height: 100vh;
background: url("${HomeImage}")
`

const StyledText = styled.h1`
margin: 0;
position: absolute;
top: calc(50% - 0.5em);
width: 100%;
text-align: center;
line-height: 1;
font-size: 8vw;
font-family: "Montserrat", sans-serif;
text-shadow: 10px 10px 0 #fff;
`

export default Home;