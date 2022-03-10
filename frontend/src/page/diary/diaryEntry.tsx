import { useEffect, useState, VFC, SyntheticEvent } from 'react';
import { MFInput } from 'component/column/input/mf_input';
import { MFPicklist } from 'component/column/input/mf_picklist';
import { MFMultiPicklist } from 'component/column/input/mf_multipicklist';
import { MFTextarea } from 'component/column/input/mf_textarea';
import { MFButton } from 'component/button/mf_button';
import {
    Column, 
    TextColumn, 
    NumberColumn, 
    PicklistColumn, 
    MultiPicklistColumn, 
    TextareaColumn, 
} from 'component/column/columnDef';
import {
    textColumns, 
    numberColumns,
    picklistColumns,
    multiPicklistColumns,
    textareaColumns
} from 'component/column/columns';
import styled from "styled-components";
import { StyledWrapper, StyledOuter } from 'commonCSS/style'
import { testData } from 'testData'
import { order } from 'page/diary/dialyColumnOrder'
import { ModalSppiner } from 'component/spinner/modalSpinner'
import { useNavigate } from 'react-router-dom'


const DiaryEntry: VFC = () => {
    // 各項目のuseStateを宣言
    const [place, setPlace] = useState('');
    const [caughtFish, setCaughtFish] = useState(['']);
    const [comment, setComment] = useState('');
    const [rod, setRod] = useState(['']);
    const [method, setMethod] = useState(['']);
    const [lure, setLure] = useState(['']);
    const [weather, setWeather] = useState('');
    const [wind, setWind] = useState(0);
    const [tide, setTide] = useState('');

    const [orderedColumns, setOrderedColumns] = useState<Column[]>();
    const [isLoad, setIsLoad] = useState<boolean>(false);

    const navigate = useNavigate();

    // URLパラメータからid値を取得する
    // 取得値がidの正規表現に合致しない場合は新規レコード作成画面とみなす

    // ページ初期化
    useEffect(()=> {       
        (
            async () => {
                console.log('useEffect start');
                // 項目定義情報の取得
                const columns: Column[] = [...textColumns, ...numberColumns, ...picklistColumns, ...multiPicklistColumns, ...textareaColumns];
            
                // ここから編集時のみの処理
                    // サーバーから初期値を取得 await

                    // サーバーからの値をrecordに注入
                    // for(let c of columns){
                    //     if(testData[c.id]) c.value = testData[c.id];
                    // }

                    // 初期値を各stateにセット
                    // setPlace(testData.place ?? '');
                    // setCaughtFish(testData.caughtFish ?? []);
                    // setComment(testData.comment ?? '');
                    // setRod(testData.rod ?? []);
                    // setMethod(testData.method ?? []);
                    // setLure(testData.lure ?? []);
                    // setWeather(testData.weather ?? '');
                    // setWind(testData.wind ?? 0);
                    // setTide(testData.tide ?? '');
                //ここまで編集時のみの処理
                
                // columnOrderの順に項目の並び替え
                injectHandleChange(columns);
                setOrderedColumns(sortColumns(order, columns));
                console.log(orderedColumns);
            })()
    },[])
        
    const sortColumns = (order:string[], columns:Column[]):Column[] => {
        const result: Column[] = [];
        
        for( let o of order ){
            for(let c of columns){
                if(c.id === o){
                    result.push(c);
                    break;
                }
            }
        }
        return result;
    }

    const injectHandleChange = (columns:Column[]):void => {
        for(let c of columns){
            if(c.id === 'place') c.handleChange = setPlace;
            if(c.id === 'caughtFish') c.handleChange = setCaughtFish;
            if(c.id === 'comment') c.handleChange = setComment;
            if(c.id === 'rod') c.handleChange = setRod;
            if(c.id === 'method') c.handleChange = setMethod;
            if(c.id === 'lure') c.handleChange = setLure;
            if(c.id === 'weather') c.handleChange = setWeather;
            if(c.id === 'wind') c.handleChange = setWind;
            if(c.id === 'tide') c.handleChange = setTide;
        }
    }
    
    const submit =async (e: SyntheticEvent)=> {
        console.log('★submit★');
        setIsLoad(true);
        e.preventDefault();
        const data = {
            place: place,
            caughtFish: caughtFish,
            comment: comment,
            rod: rod,
            method: method,
            lure: lure,
            weather: weather,
            wind: wind,
            tide: tide
        }
        console.log('data');
        console.log(data);
        navigate('/view', { state: { id:1 }});
    }


    return (
        <StyledOuter>
            {isLoad && <ModalSppiner height={100} width={100} />}
            <StyledWrapper>
                <form method="post" onSubmit={ submit } >
                    <StyledContainer>
                            {orderedColumns?.map((oc)=>{
                                if(isInput(oc)){
                                    console.log(oc.value);
                                    return <MFInput
                                        label={oc.label}
                                        id={oc.id}
                                        input={oc.input}
                                        key={oc.id}
                                        value={oc.value}
                                        handleChange={oc.handleChange}></MFInput>
                                }else if(isPickList(oc)){
                                    console.log(oc.value);
                                    return <MFPicklist
                                        label={oc.label}
                                        id={oc.id}
                                        select={oc.select}
                                        options={oc.options}
                                        key={oc.id}
                                        value={oc.value}
                                        handleChange={oc.handleChange}></MFPicklist>
                                }else if(isMultiPickList(oc)){
                                    console.log(oc.value);
                                    return <MFMultiPicklist
                                        label={oc.label}
                                        id={oc.id}
                                        select={oc.select}
                                        options={oc.options}
                                        key={oc.id}
                                        value={oc.value}
                                        handleChange={oc.handleChange}></MFMultiPicklist>
                                }else if(isTextarea(oc)){
                                    return <MFTextarea
                                        label={oc.label}
                                        id={oc.id}
                                        textarea={oc.textarea}
                                        key={oc.id}
                                        value={oc.value}
                                        handleChange={oc.handleChange}></MFTextarea>
                                }
                            })}
                    </StyledContainer>
                <StyledButtonWrapper>
                    <StyledButtonContainer>
                        <MFButton type={'submit'} disabled={false} mode={'positive'} hidden={false} >保存</MFButton>
                    </StyledButtonContainer>
                    <StyledButtonContainer>
                        <MFButton type={'button'} disabled={false} mode={'negative'} hidden={false} >戻る</MFButton>
                    </StyledButtonContainer>
                </StyledButtonWrapper>
                </form>
            </StyledWrapper>
        </StyledOuter>
    )
}


const isInput = (column: Column): column is (TextColumn | NumberColumn) => column.type === 'input';
const isPickList = (column: Column): column is PicklistColumn => column.type === 'picklist';
const isMultiPickList = (column: Column): column is MultiPicklistColumn => column.type === 'multipicklist';
const isTextarea = (column: Column): column is TextareaColumn => column.type === 'textarea';

const StyledContainer =styled.div`
display: grid;
padding-top: 100px;
grid-auto-rows: 150px;
grid-template-columns: 50% 50%;
`;

const StyledButtonWrapper =styled.div`
width: 30%;
margin: 0 auto;
height: 150px;
align-items: center;
justify-content: center;
display: flex;
z-index: 0;
`;

const StyledButtonContainer =styled.div`
margin: 0 10%;
`;

export default DiaryEntry;