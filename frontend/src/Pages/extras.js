const DisplayScanResponse = ({ responseData }) => {
  const renderObject = (obj) => {
    if (obj === null) {
      return <span>null</span>;
    }

    return (
      <ul>
        <li>{obj.Name}</li>
        {Object.entries(obj).map(([key, value]) => (
          <li key={key}>
            <strong>{key}:</strong> {renderValue(value)}
          </li>
        ))}
      </ul>
    );
  };

  const renderArray = (arr) => {
    return (
      <ul>
        {arr.map((item, index) => (
          <li key={index}>{renderValue(item)}</li>
        ))}
      </ul>
    );
  };

  const renderValue = (value) => {
    if (typeof value === "object") {
      if (Array.isArray(value)) {
        return renderArray(value);
      } else {
        return renderObject(value);
      }
    } else {
      return String(value);
    }
  };

  return (
    <div>
      <h2>Response Data:</h2>
      {renderObject(responseData)}
    </div>
  );
}

