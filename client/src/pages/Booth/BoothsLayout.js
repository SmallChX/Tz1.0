import React, { useState, useEffect } from "react";
import './booth_layout.css'

function BoothsLayout({ selectedBooths, booths, ownedBooths, deselectedBooths, toggleBoothSelection, action }) {
    const [classBorder, setClassBorder] = useState("");

    useEffect(() => {
        switch (action) {
            case 'register':
                setClassBorder("border-add");
                break;
            case 'change':
                setClassBorder("border-edit");
                break;
            case 'remove':
               setClassBorder("border-delete");
                break;
            default:
                setClassBorder("");
                break;
            }
    }, [action]);

    const getPriceClass = (booth) => {
        if (booth.price === 12000000) {
            return 'cheap-booth';
        } else if (booth.price === 15000000) {
            return 'expensive-booth';
        }
        return '';
    };

    const isBoothHidden = (booth) => {
        if (ownedBooths.length > 0) {
            const isOwned = ownedBooths.includes(booth.ID);
            if (isOwned) {
                return false;
            }
        }
        return (booth.company_info.company_id && booth.company_info.company_id !== 0) || booth.Level === 2;
    };

    

    return (
        <div className="booth-layout-container booth-layoutcolumn booth-layoutmiddle justify-content-center">
            <div className={classBorder} style={{backgroundColor:"white", borderRadius:"15px"}}>
                <div className="booth-layoutH6 booth-layouttop d-flex justify-content-center align-items-center" style={{borderTopLeftRadius:"15px", borderTopRightRadius:"15px"}}>
                    <span style={{fontSize:"24px",fontWeight:"600",color:"white", textAlign: "center"}}>Tòa nhà BK.B6 - Trường Đại học Bách khoa - ĐHQG-HCM</span>
                </div>
            <div className="booth-layoutrow">
            <div className="booth-layoutside-A booth-layoutcolumn">
                <div className="booth-layoutblock-1 booth-layoutrow">
                    <div>
                        <div className="booth-layoutH6-side booth-layoutrow">
                            <div className="booth-layoutH6">
                            </div>
                            <div className="booth-layoutH6-side booth-layoutcolumn">
                                <div className="booth-layoutgrass plaque" style={{margin:"10px", width:"100px", height:"50px"}}></div>
                            </div>
                        </div>
                        <div className="booth-layoutleft">
                            <div className="booth-layoutbooth-1 booth-layoutrow booth-layoutmiddle" style={{marginTop:"30px"}}>
                                <div className="booth-layoutcolumn booth-layoutblock-a">
                                        {booths.slice(0, 2).map(booth => (
                                            <div key={booth.ID} 
                                            className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                            onClick={() => toggleBoothSelection(booth.ID)}    
                                            ><span>{booth.ID}</span></div>
                                        ))}
                                </div>
                                <div className="booth-layoutcolumn booth-layoutspace booth-layoutblock-b">
                                        <div className="booth-layoutrow booth-layoutspace">
                                            {booths.slice(2, 4).map(booth => (
                                                <div key={booth.ID} 
                                                className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}
                                                        ${getPriceClass(booth)}`} 
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                            ))}
                                        </div>
                                        <div className="booth-layoutrow booth-layoutspace">
                                            {booths.slice(4, 6).map(booth => (
                                                <div key={booth.ID} 
                                                className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}
                                                        ${getPriceClass(booth)}`} 
                                                onClick={() => toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                            ))}
                                        </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="booth-layoutright booth-layoutcolumn booth-layoutmiddle">
                        <div className="booth-layoutH6 stair" style={{width:"80px", height:"20px"}}>
                        </div>
                        <div className="booth-layoutstage booth-layoutentrance d-flex justify-contents-center align-items-center" style={{height:"60px"}}>
                            <span style={{fontSize:"20px", color:"white", fontWeight:"600", textAlign: "center"}}>Sân khấu</span>
                        </div>
                        <div className="booth-layoutmiddle-grass booth-layoutrow booth-layoutspace">
                            <div className="booth-layoutgrass plaque" style={{width:"40px", height:"30px"}}></div>
                            <div className="booth-layoutgrass plaque" style={{width:"40px", height:"30px"}}>.</div>
                        </div>
                    </div>
                </div>
                
            </div>
            <div className="booth-layoutside-B booth-layoutcolumn">
                <div className="booth-layoutblock-1 booth-layoutcolumn">
                    <div className="booth-layoutrow booth-layoutpadding">
                        <div className="booth-layoutrow">
                            {booths.slice(6, 10).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(10, 14).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(14, 18).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                    </div>
                    <div className="booth-layoutrow booth-layoutpadding">
                        <div className="booth-layoutrow">
                            {booths.slice(18, 20).map(booth => (
                                <div key={booth.ID} 
                                className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}
                                                        ${getPriceClass(booth)}`} 
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(20, 24).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(24, 28).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                    </div>
                    <div style={{position:"relative"}}>
                        <div className="booth-layoutgrass long-plaque" style={{height:"34px"}}></div>
                        <div style={{position:"absolute", top:"0",left:"50px", width:"500px", height:"33px", backgroundColor:"#1a4d2e"}}></div>
                    </div>
                </div>
                <div className="booth-layoutblock-2 booth-layoutcolumn">
                    <div className="booth-layoutrow booth-layoutpadding">
                        <div className="booth-layoutrow">
                            {booths.slice(28, 32).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(32, 36).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(36, 40).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                    </div>
                    <div className="booth-layoutrow booth-layoutpadding">
                        <div className="booth-layoutrow">
                            {booths.slice(40, 44).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(44, 48).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                        <div className="booth-layoutrow">
                            {booths.slice(48, 52).map(booth => (
                                <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                            onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
            </div>
            <div className="booth-layoutbottom booth-layoutrow">
                <div>
                    <div className="booth-layoutrow ">
                        <div className="booth-layoutgrass plaque" style={{height:"34px", width:"120px"}}></div>
                        <div className="booth-layoutentrance booth-layoutbackdrop">
                                    Backdrop
                                </div>
                    </div>
                    <div></div>
                    <div className="booth-layoutblock-3 booth-layoutrow booth-layoutspace">
                    
                    <div className="booth-layoutpadding">
                        <div className="booth-layoutentrance booth-layoutgate d-flex justify-contents-center align-items-center">
                        <span style={{fontWeight:"500", fontSize:"16px", color:"white", textAlign: "center"}}>
                            Cổng chào
                            </span>
                        </div>
                    </div>
                    <div className="booth-layoutcolumn booth-layoutspace booth-layoutcenter">
                            <div className="booth-layoutspace">
                                <div className="booth-layoutentrance booth-layoutcheckin d-flex justify-contents-center align-items-center mx-3" style={{width:"60px", height:"40px"}}>
                                    <span style={{fontWeight:"500", fontSize:"16px", color:"white", textAlign: "center"}}>Checkin</span>
                                </div>

                            </div>
                            <div className="booth-layoutentrance booth-layoutglory d-flex justify-contents-center align-items-center mt-2" 
                                    style={{width:"140px", height:"70px"}}>
                                <span style={{fontWeight:"500", fontSize:"16px", color:"white", textAlign: "center"}}>
                                    Không gian tuyên dương Sinh viên tiêu biểu

                                </span>
                            </div>
                        
                    </div>
                </div>
                </div>
                
                <div className="booth-layoutcolumn booth-layoutblock-bottom-booth ">
                <div style={{position:"relative"}}>
                        <div className="booth-layoutgrass long-plaque" style={{height:"34px"}}></div>
                        <div style={{position:"absolute", top:"0",left:"50px", width:"500px", height:"34px", backgroundColor:"#1a4d2e"}}></div>
                    </div>
                    <div className="booth-layoutrow booth-layoutpadding">
                            <div className="booth-layoutrow">
                                {booths.slice(52, 56).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                            <div className="booth-layoutrow">
                                {booths.slice(56, 60).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                            <div className="booth-layoutrow">
                                {booths.slice(60, 64).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                        </div>
                        <div className="booth-layoutrow booth-layoutpadding">
                            <div className="booth-layoutrow">
                                {booths.slice(64, 68).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                            <div className="booth-layoutrow">
                                {booths.slice(68, 72).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                            <div className="booth-layoutrow">
                                {booths.slice(72, 76).map(booth => (
                                    <div key={booth.ID} className={`booth-layoutindividual-booth 
                                                        ${selectedBooths.includes(booth.ID) ? 'booth-layoutselected' : ''}
                                                        ${deselectedBooths.includes(booth.ID) ? 'booth-layoutdeselected' : ''}
                                                        ${ownedBooths.includes(booth.ID) ? 'booth-layoutowned' : ''}
                                                        ${getPriceClass(booth)} 
                                                        ${isBoothHidden(booth) ? 'booth-layouthidden' : ''}`}
                                onClick={() => !isBoothHidden(booth) && toggleBoothSelection(booth.ID)}><span>{booth.ID}</span></div>
                                ))}
                            </div>
                        </div>
                    </div>
            </div>
            </div>
        </div>  
    )
}

export default BoothsLayout