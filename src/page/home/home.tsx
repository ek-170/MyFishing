import { VFC } from 'react';
import { StyledWrapper, StyledOuter } from 'commonCSS/style'

const Home: VFC = () => {
    return (
        <StyledOuter>
            <StyledWrapper>
                <h1>Here is Home</h1>
            </StyledWrapper>
        </StyledOuter>
    )
}

export default Home;