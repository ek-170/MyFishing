import { useEffect, useState, VFC } from 'react';
import { MFButton } from 'component/button/mf_button';
import { MFOutoPut } from 'component/column/output/mf_output'
import styled from "styled-components";
import { StyledWrapper, StyledOuter } from 'commonCSS/style'
import { testData } from 'testData'
import { order } from 'page/diary/dialyColumnOrder'
import { Column } from 'component/column/columnDef';
import {
    textColumns, 
    numberColumns,
    picklistColumns,
    multiPicklistColumns,
    textareaColumns
} from 'component/column/columns';

const DiaryView: VFC = () => {
    
    const [orderedColumns, setOrderedColumns] = useState<Column[]>();

    // URLパラメータからid値を取得する
    // 取得値がidの正規表現に合致しない場合は新規レコード作成画面とみなす
    
    // ページ初期化
    useEffect(()=> {       
        (
            async () => {
                console.log('@Diary View@');
                console.log('useEffect start');
                // 項目定義情報の取得
                const columns: Column[] = [...textColumns, ...numberColumns, ...picklistColumns, ...multiPicklistColumns, ...textareaColumns];
                console.log('columns');
                console.log(columns);
                
                // サーバーから初期値を取得

                // サーバーからの値をrecordに注入
                for(let c of columns){
                    if(testData[c.id]) c.value = testData[c.id];
                }

                setOrderedColumns(sortColumns(order, columns));
                console.log('orderedColumns');
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


    return (
        <StyledOuter>
            <StyledWrapper>
                <StyledContainer>
                    {orderedColumns?.map((oc)=>{
                        return <MFOutoPut
                                label={oc.label}
                                id={oc.id}
                                key={oc.id}
                                value={oc.value}></MFOutoPut>
                        })}
                </StyledContainer>
                <StyledButtonWrapper>
                    <StyledButtonContainer>
                        <MFButton type={'button'} disabled={false} mode={'positive'}>編集</MFButton>
                    </StyledButtonContainer>
                    <StyledButtonContainer>
                        <MFButton type={'button'} disabled={false} mode={'negative'} hidden={true}>戻る</MFButton>
                    </StyledButtonContainer>
                </StyledButtonWrapper>
            </StyledWrapper>
        </StyledOuter>
    )
}

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
`;

const StyledButtonContainer =styled.div`
margin: 0 10%;
`;

export default DiaryView;