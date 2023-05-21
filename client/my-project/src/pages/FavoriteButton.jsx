import React, { useState } from 'react';
import {GrFavorite} from 'react-icons/gr'
import {MdFavorite} from 'react-icons/md'

const FavoriteButton = () => {
  const [isFavorite, setIsFavorite] = useState(false);

  const Favorite = () => {
    return (
        <>
            <button><MdFavorite className='text-red-700'/></button>
        </>
    )
  }

  const DeleteFavorite = () => {
    return (
        <>
            <button><GrFavorite/></button>
        </>
    )
  }

  const handleFavoriteClick = () => {
    setIsFavorite(!isFavorite);
  };

  return (
    <button onClick={handleFavoriteClick}>
      {isFavorite ? <Favorite/> : <DeleteFavorite/>}
    </button>
  );
};

export default FavoriteButton;