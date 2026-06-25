interface ProfileCardProps{
    nama: string;
    role: string;
}

const ProfileCard: React.FC<ProfileCardProps> = ({
    nama,
    role,
}) => {
    return (
        <div className="profile-card">
            <h2>{nama}</h2>
            <p>{role}</p>
        </div>
    )
}
export default ProfileCard;