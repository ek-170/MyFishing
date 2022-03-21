import { VFC, useState, useRef, useEffect} from 'react';
import styled from "styled-components";
import GlassImage from 'img/GlassImage.png';
import PenImage from 'img/PenImage.png';
import HomeImage from 'img/HomeImage.png';
import GearImage from 'img/GearImage.png';
import { IconLink } from 'component/header/iconLink';
import { container } from 'utils/const/size';
import character from 'utils/const/character';

const Header: VFC = () => {
    const [isHidden, setIsHidden] = useState(true);
    const menuRef = useRef<HTMLUListElement>(null);
    const handleClickDocument = useRef((e:Event)=>{});

    useEffect(() => {
        handleClickDocument.current = (e:Event) => {
            if(!(e.target instanceof HTMLElement)) return;
            if(!menuRef.current?.contains(e.target)){
                setIsHidden(true);
                document.removeEventListener('click', handleClickDocument.current);
            }
        }
    },[])

    useEffect(() => {
        isHidden === false && document.addEventListener('click', handleClickDocument.current);
    },[isHidden]);

    const handleMenuClick = ()=> {
        setIsHidden(!isHidden);
    }

    const headerIconLinks =[{
        id: 'homeIcon',
        to: '/',
        src: HomeImage,
        alt: 'home',
        width: '40px',
    },{
        id: 'editiIcon',
        to: '/entry/new',
        src: PenImage,
        alt: 'edit',
        width: '40px',
    },{
        id: 'searchIcon',
        to: '/search',
        src: GlassImage,
        alt: 'search',
        width: '40px',
    }]

    const dummyConfigLinks = [ 1,2,3 ]
    
    return (
        <StyledWrapper>
            <StyledContainer>
                <StyledFlex>
                    {headerIconLinks.map((headerIconLink)=>(
                        <IconLink
                            to={headerIconLink.to}
                            src={headerIconLink.src}
                            alt={headerIconLink.alt}
                            width={headerIconLink.width}
                            key={headerIconLink.id}
                        ></IconLink>
                    ))}
                        <StyledConfig onClick={()=>{handleMenuClick()}}>
                            <StyledConfigImg src={GearImage} alt='Setting Menu'></StyledConfigImg>
                            <StyledConfigLinkLists ref={menuRef} hidden={isHidden}>
                                {dummyConfigLinks.map((index)=>(
                                    <StyledConfigLinkList key={index}><StyledConfigLink>menu{index}</StyledConfigLink></StyledConfigLinkList>
                                ))}
                            </StyledConfigLinkLists>
                        </StyledConfig>
                </StyledFlex>
            </StyledContainer>
        </StyledWrapper>
    )
}

const StyledWrapper =styled.div`
    position: fixed;
    z-index: 100;
    height: 60px;
    background: black;
    top:0;
    left: 0;
    right: 0;
`

const StyledContainer =styled.div`
    width: ${container.WIDTH};
    margin-left: auto;
    margin-right: auto;
`

const StyledFlex =styled.div`
    display: flex;
    positon: relative;
`
const StyledConfig =styled.div`
    margin-left: auto;
    position: relative;
    height: 40px;
    width: 40px;
    `
    
    const StyledConfigImg =styled.img`
    width:40px;
    height:auto;
    padding: 8px 16px;
    transition: 0.2s ;
    cursor:pointer;
&:hover{
    background: #383d43;
}
`

const StyledConfigLinkLists =styled.ul`
    position: absolute;
    list-style: none;
    top: 100%;
    right: 0;
`
const StyledConfigLinkList =styled.li`
    font-family: ${character.FONT};
    background: white;
    height: 50px;
    line-height: 50px;
    padding:0px 15px 0px 15px;
    display: inline-block;
    width: auto;
    text-align: center;
    white-space: nowrap;
    border: 0.5px solid;
    border-color: #778899;
`

const StyledConfigLink =styled.a`
    display: inline-block;
    vertical-align: middle;
`


export default Header