function Prompt() {
  let toast = function (c) {
    const { title = "", icon = "success", position = "top-end" } = c;
    const Toast = Swal.mixin({
      toast: true,
      title,
      position,
      icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({});
  };

  let success = function (c) {
    const { title = "", text = "", footer = "" } = c;
    Swal.fire({
      icon: "success",
      title,
      text,
      footer,
    });
  };

  let error = function (c) {
    const { title = "", text = "", footer = "" } = c;
    Swal.fire({
      icon: "error",
      title,
      text,
      footer,
    });
  };

  async function custom(c) {
    const { icon = "", title = "", html = "", showConfirmButton = true } = c;

    const { value: result } = await Swal.fire({
      icon,
      title,
      html,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      preConfirm: () => {
        return [
          document.getElementById("start").value,
          document.getElementById("end").value,
        ];
      },
      didOpen: () => {
        if (c.didOpen !== undefined) {
          c.didOpen();
        }
      },
    });

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  }
  return {
    toast,
    success,
    error,
    custom,
  };
}
