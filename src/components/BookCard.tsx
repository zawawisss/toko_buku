import { useState } from "react";


interface BookCardProps {
    id: number;
    judul: string;
    penulis: string;
    harga: number;
    stok: number;
    onDelete: (id: number) => void
    onBuy: (id: number) => void
}

function BookCard({id, judul, penulis, harga, stok, onDelete, onBuy}: BookCardProps){
    return(
        <div className="book-card">
            <h3>{judul}</h3>
            <p>Penulis : {penulis}</p>
            <p>Harga : {harga}</p>
            <p>Stok : {stok}</p>
            <button disabled={stok === 0} onClick={() => onBuy(id)}>{stok > 0 ? "Beli" : "Habis"}</button>
            <button onClick={() => onDelete(id)} style={{backgroundColor: "#e74c3c", color: "white", marginLeft: "10px"}}>Hapus</button>
        </div>
    )
}
export default BookCard