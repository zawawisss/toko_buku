import BookCard from "./components/BookCard";
import Header from "./components/Header";
import ProfileCard from "./components/ProfileCard";
import "./App.css"
import React, { useEffect, useState } from "react";

interface Buku{
    id: number;
    judul: string;
    penulis: string;
    harga: number;
    stok: number;
}

function App(){
    const [daftarBuku, setDaftarBuku] = useState<Buku[]>([])
    const [judul, setJudul] = useState("")
    const [penulis, setPenulis] = useState("")
    const [harga, setHarga] = useState("")
    const [stok, setStok] = useState("")
    const handleSubmit = (e: React.SubmitEvent) => {
        e.preventDefault()
        const bukuBaru = {
            judul,
            penulis,
            harga: Number(harga),
            stok: Number(stok)
        }
        fetch("http://localhost:8080/books", {
            method: "POST",
            headers:{"Content-Type" : "application/json"},
            body: JSON.stringify(bukuBaru)
        })
        .then((res) => {
            if (res.status === 201) {
                return res.json()
            }
            throw new Error("Gagal menambah buku")
        })
        .then((data: Buku) => {
            setDaftarBuku([...daftarBuku, data])

            setJudul("")
            setPenulis("")
            setHarga("")
            setStok("")
        })
        .catch((error) => console.error(error))
    }
    const handleBuy = (id:number) => {
        const buku = daftarBuku.find((b) => b.id === id);
        if (!buku || buku.stok === 0) return
        const bukuUpdate = {...buku, stok: buku.stok - 1}
        fetch("http://localhost:8080/books/" + id,{
            method: "PUT",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(bukuUpdate)
        })
        .then((res) => {
            if (res.ok){
                const bukuLama = daftarBuku.map((b) => b.id === id ? bukuUpdate : b)
                setDaftarBuku(bukuLama)
            }
        })
    }
    const handleDelete = (id:number) => {
        fetch("http://localhost:8080/books/" + id, {
            method: "DELETE"
        })
        .then((res) => {
            if (res.ok) {
                const sisaBuku = daftarBuku.filter((buku) => buku.id !== id);
                setDaftarBuku(sisaBuku)
            }
        })
    }
    useEffect(() => {
        fetch("http://localhost:8080/books")
        .then((response) => response.json())
        .then((data) => setDaftarBuku(data))
    }, [])
    return(
        <>
            <Header/>
            <ProfileCard nama="Zawawi" role="Programmer"/>
            <form onSubmit={handleSubmit} className="form-tambah-buku">
                <h3>Tambah Buku Baru</h3>
                <div>
                    <input 
                        type="text"
                        placeholder="Judul Buku"
                        value={judul}
                        onChange={(e) => setJudul(e.target.value)}
                        required
                    />
                    <input 
                        type="text"
                        placeholder="Penulis"
                        value={penulis}
                        onChange={(e) => setPenulis(e.target.value)}
                        required
                    />
                    <input 
                        type="text"
                        placeholder="Harga"
                        value={harga}
                        onChange={(e) => setHarga(e.target.value)}
                        required
                    />
                    <input 
                        type="text"
                        placeholder="Stok"
                        value={stok}
                        onChange={(e) => setStok(e.target.value)}
                        required
                    />
                    <button type="submit">Tambah Buku</button>
                </div>
            </form>
            <div className="container-project">
                {daftarBuku.map((buku) => (
                    <BookCard
                        key={buku.id}
                        id={buku.id}
                        judul={buku.judul}
                        penulis={buku.penulis}
                        harga={buku.harga}
                        stok={buku.stok}
                        onBuy={handleBuy}
                        onDelete={handleDelete}
                    />  
                ))}
            </div>
        </>
    )
}
export default App